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
	fmt.Printf("攻击：")
	fmt.Println(role.Attack)
	fmt.Printf("防御：")
	fmt.Println(role.Defense)
	fmt.Printf("暴击：")
	fmt.Println(role.ViolentRate)
	fmt.Printf("爆伤：")
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
func GetRole(cookie string) model.RoleStauts {
	//urlPath := "https://dao.fushisanlang.cn/role/get"
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
	if message.Code != 200 && message.Code != 0 {
		tools.Clean()

		fmt.Println(message.Message)

		AddRole(cookie)
	} else {

		json.Unmarshal([]byte(respBody), &Role)

		//return serverVersion.Generation, serverVersion.WorldMp
	}
	return Role
}

func addRole(cookie string) model.RoleCreate {
	var RoleCreate model.RoleCreate
	var name string
	fmt.Println("请输入角色名")
	fmt.Scan(&name)
	RoleCreate.Name = name
	var sex string
	sexCode := 3
	for true {
		tools.Clean()
		fmt.Println("请输入性别， 男 或 女 ")
		fmt.Scan(&sex)

		if sex == "男" {
			sexCode = 1
			break
		} else if sex == "女" {
			sexCode = 0
			break
		} else {
			tools.Clean()
			fmt.Println("输入错误")

		}
	}

	RoleCreate.Sex = sexCode
	var j, m, s, h, t int
	var randCode string
	i := 1
	for true {
		j, m, s, h, t = tools.GetRoleStableStatus()
		tools.Clean()
		fmt.Println("基础属性为：金 " + strconv.Itoa(j) + ", 木 " + strconv.Itoa(m) + ", 水 " + strconv.Itoa(s) + ", 火 " + strconv.Itoa(h) + ", 土 " + strconv.Itoa(t))
		if i == 3 {
			fmt.Println("随即次数用尽")
			break
		}
		fmt.Println("基础属性随机生成，将影响最终角色各项属性，如不满意，可以输入")
		fmt.Println("剩余可随机次数 " + strconv.Itoa(3-i) + " 次,重新生成还是下一步？")
		fmt.Println("输入 继续 进行下一步，输入 重选 重新生成")

		fmt.Scan(&randCode)
		if randCode == "继续" {
			break
		} else if randCode == "再来一次" {
			i = i + 1
		} else {
			fmt.Println("输入异常，使用此属性")
			time.Sleep(3 * time.Second)
			break
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
	for true {

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

		//RoleCreate.Jin = RoleCreate.Jin + schoolInfo.Jin
		//RoleCreate.Tu = RoleCreate.Tu + schoolInfo.Tu
		//RoleCreate.Mu = RoleCreate.Mu + schoolInfo.Mu
		//RoleCreate.Shui = RoleCreate.Shui + schoolInfo.Shui
		//RoleCreate.Huo = RoleCreate.Huo + schoolInfo.Huo

		fmt.Println("将选择 " + schoolInfo.Name + "作为起始门派")
		fmt.Println("输入 继续 进行下一步，输入 重选 返回门派选择")
		var schoolStatus string
		fmt.Scan(&schoolStatus)
		if schoolStatus == "继续" {
			RoleCreate.School_id = schoolId
			//fmt.Println(RoleCreate.Name)
			break
		} else if schoolStatus == "重选" {
			tools.Clean()
		} else {
			tools.Clean()
			fmt.Println("输入异常，重新选择门派")
		}

	}

	fmt.Println("log 1")
	SkillList := GetSkillList(cookie)
	//fmt.Println(len(SkillList))
	fmt.Println("log 2")
	tools.Clean()
	var skillInfo model.SkillInfo
	for true {
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
		fmt.Println("输入 继续 进行下一步，输入 重选 返回技能选择")
		var skillStatus string
		fmt.Scan(&skillStatus)
		if skillStatus == "继续" {
			RoleCreate.Skill_id = skillId
			break
		} else if skillStatus == "重选" {
			tools.Clean()
		} else {
			tools.Clean()
			fmt.Println("输入异常，重新选择技能")
		}
	}

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
