/*
@Time : 2021/9/7 13:16
@Author : zhangyin
@File : role
@Software: GoLand
*/
package service

import (
	"Dao-client/app/model"
	"Dao-client/app/tools"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func PrintRole(role model.RoleStauts) {
	fmt.Printf("角色名：")
	fmt.Println(role.Name)
	fmt.Printf("性别：")
	if role.Sex == 0 {
		fmt.Println("女")
	} else {
		fmt.Println("男")
	}
	fmt.Printf("等级：")
	fmt.Println(role.Level)
	fmt.Printf("经验：")
	fmt.Println(role.Experience)
	fmt.Printf("门派：")
	fmt.Println(role.School)
	fmt.Printf("功法：")
	fmt.Println(role.Skill)
	fmt.Printf("武器：")
	fmt.Println(role.Arm)
	fmt.Printf("攻击：")
	fmt.Println(role.Attack)
	fmt.Printf("防御：")
	fmt.Println(role.Defense)
	fmt.Printf("暴击：")
	fmt.Println(role.ViolentRate)
	fmt.Printf("暴伤：")
	fmt.Println(role.ViolentAttack)
	fmt.Printf("冷却：")
	fmt.Println(role.SkillCD)

	fmt.Printf("元气：")
	fmt.Printf("%d / %d \n", role.Hp, role.MaxHp)
	fmt.Printf("精气：")
	fmt.Printf("%d / %d \n", role.Mp, role.MaxMp)
	fmt.Printf("回元：")
	fmt.Println(role.HpBak)

	fmt.Printf("聚气：")
	fmt.Println(role.MpBak)

}
func GetRole(cookie string) (model.RoleStauts, bool) {

	urlPath := UrlPre + "/role/get"

	postString := "{}"
	postStringByte := []byte(postString)
	req, err := http.NewRequest("GET", urlPath, bytes.NewBuffer(postStringByte))
	req.Header.Set("Cookie", cookie)
	client := &http.Client{}
	resp, _ := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	respBody := string(body)

	defer resp.Body.Close()

	var message model.RespMessage
	json.Unmarshal([]byte(respBody), &message)
	var Role model.RoleStauts
	var status bool
	if message.Code != 200 && message.Code != 0 {
		tools.Clean()

		fmt.Println(message.Message)

		AddRole(cookie)
		status = false
	} else {

		json.Unmarshal([]byte(respBody), &Role)
		status = true

		//return serverVersion.Generation, serverVersion.WorldMp
	}
	return Role, status
}

func addRole(cookie string) model.RoleCreate {
	var RoleCreate model.RoleCreate
	var name string
	fmt.Println("请输入角色名")
	fmt.Scan(&name)
	RoleCreate.Name = name
	var sex int

	tools.Clean()
	fmt.Println("请选择性别：\n1. 男\n2. 女 ")
	fmt.Scan(&sex)

	if sex == 1 {
		sex = 1
	} else if sex == 2 {
		sex = 0
	} else {
		tools.Clean()
		fmt.Println("输入错误,使用男性作为性别。")
		sex = 1

	}

	RoleCreate.Sex = sex
	var j, m, s, h, t, randCode int

	for i := 0; i < 4; i++ {
		fmt.Println("===========")
		fmt.Println(i)
		fmt.Println("===========")
		j, m, s, h, t = tools.GetRoleStableStatus()
		tools.Clean()
		fmt.Println("基础属性为：金 " + strconv.Itoa(j) + ", 木 " + strconv.Itoa(m) + ", 水 " + strconv.Itoa(s) + ", 火 " + strconv.Itoa(h) + ", 土 " + strconv.Itoa(t))

		fmt.Println("基础属性随机生成，将影响最终角色各项属性，如不满意，可以输入")
		if i < 3 {
			fmt.Println("剩余可随机次数 " + strconv.Itoa(3-i) + " 次,重新生成还是下一步？")
			fmt.Println("1. 继续\n2. 重选 ")
			fmt.Println("输入 1 或者 2 选择操作")
			fmt.Scan(&randCode)
			if randCode == 1 {
				break
			} else if randCode == 2 {

			} else {
				fmt.Println("输入异常，使用此属性")

				break
			}
		} else {
			fmt.Println("随机次数用尽，使用此属性")
		}

	}
	tools.Clean()
	fmt.Println("最终属性为：金 " + strconv.Itoa(j) + ", 木 " + strconv.Itoa(m) + ", 水 " + strconv.Itoa(s) + ", 火 " + strconv.Itoa(h) + ", 土 " + strconv.Itoa(t))
	RoleCreate.Jin = j
	RoleCreate.Mu = m
	RoleCreate.Shui = s
	RoleCreate.Huo = h
	RoleCreate.Tu = t
	SchoolList := GetSchoolList(cookie)
	var schoolInfo model.SchoolInfo

	for i := 0; i < len(SchoolList); i++ {
		json.Unmarshal([]byte(SchoolList[i]), &schoolInfo)
		fmt.Printf("%d", schoolInfo.Id)
		fmt.Println(". " + schoolInfo.Name + " 简介:" + schoolInfo.Info)
	}
	fmt.Println("请输入门派前的序号选择门派")
	var schoolId int
	fmt.Scan(&schoolId)
	//fmt.Println(schoolId)

	if schoolId == 0 || schoolId > len(SchoolList) {
		fmt.Println("输入异常，使用默认门派")
		schoolId = 1
	}
	json.Unmarshal([]byte(SchoolList[schoolId-1]), &schoolInfo)

	fmt.Println("将选择 " + schoolInfo.Name + "作为起始门派")
	RoleCreate.School_id = schoolId
	SkillList := GetSkillList(cookie)

	tools.Clean()
	var skillInfo model.SkillInfo

	for i := 0; i < len(SkillList); i++ {
		json.Unmarshal([]byte(SkillList[i]), &skillInfo)

		fmt.Printf("%d", skillInfo.Id)
		fmt.Println(". " + skillInfo.Name + " 简介:" + skillInfo.Description)
	}
	fmt.Println("请输入技能前的序号选择技能")
	var skillId int
	fmt.Scan(&skillId)
	if skillId == 0 || skillId > len(SkillList) {
		skillId = 1
		fmt.Println("输入异常，使用默认技能")
	}
	json.Unmarshal([]byte(SkillList[skillId-1]), &skillInfo)
	//tools.Clean()
	fmt.Println("将选择 " + skillInfo.Name + "作为起始技能")
	RoleCreate.Skill_id = skillId
	ArmList := GetArmList(cookie)

	tools.Clean()
	var armInfo model.ArmInfo

	for i := 0; i < len(ArmList); i++ {
		json.Unmarshal([]byte(ArmList[i]), &armInfo)

		fmt.Println(fmt.Sprint(armInfo.Id) + ". " + armInfo.Name + " 简介:" + armInfo.Description)
	}
	fmt.Println("请输入武器前的序号选择武器。\n请注意，武器与功法匹配时，威力会成倍提升。")
	var armId int
	fmt.Scan(&armId)
	if armId == 0 || armId > len(ArmList) {
		armId = 1
		fmt.Println("输入异常，使用默认武器")
	}
	json.Unmarshal([]byte(ArmList[armId-1]), &armInfo)
	//tools.Clean()
	fmt.Println("将选择 " + armInfo.Name + "作为起始武器")
	RoleCreate.Arm_id = armId
	//panic("经历")
	return RoleCreate
}

func AddRole(cookie string) {
	Role := addRole(cookie)

	for true {
		urlPath := UrlPre + "/role/add"

		//postString := "{\"user\":\"" + user + "\",\"pass\":\"" + pass + "\",\"passre\":\"" + passre + "\",\"mail\":\"" + email + "\"}"
		RoleJson, _ := json.Marshal(Role)

		//postStringByte := []byte(postString)
		req, err := http.NewRequest("POST", urlPath, bytes.NewBuffer(RoleJson))

		req.Header.Set("Cookie", cookie)
		client := &http.Client{}
		resp, _ := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		respStr := string(body)
		var respMessage model.RespMessage
		json.Unmarshal([]byte(respStr), &respMessage)
		//
		if respMessage.Code == 401 {
			//有角色
			fmt.Println(respMessage.Message)
			break
		} else if respMessage.Code == 402 {
			//角色名冲突
			tools.Clean()
			fmt.Println(respMessage.Message)
			fmt.Println("请重新输入角色名")
			var name string
			fmt.Scan(&name)
			Role.Name = name
		} else if respMessage.Code == 200 {
			fmt.Println(respMessage.Message)

			break
		} else {
			tools.Clean()
			fmt.Println(respMessage.Message)
			fmt.Println("异常错误，请尝试重启或联系客服。\n客服信息： https://dao.fushisanlang.cn/address")
			fmt.Println("请重启，20秒后将自动关闭")
			time.Sleep(20 * time.Second)
			panic("自动退出")
		}

	}

}
