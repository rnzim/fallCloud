package docker

import (
	
	"fmt"
	"os/exec"
	"strings"
	
)


func CreateContainer(idService int) (string ,string){ 
	if idService == 0{
		result ,idContainer := FireFox()

		if len(result) > 1{
			return result,idContainer
		}
	}else if idService == 1{
		result ,idContainer := Minecraft()

		if len(result) > 1{
			return result,idContainer
		}
	}
	
	return "Serviço Invalido!","Null"

}
func StartContainerByID(containerID string) string  {
	cmd,err := exec.Command("docker","start",containerID).Output()
	if err!= nil{
		panic(err.Error())
		
	}
	out_str := string(cmd)

	return out_str

}

func ListAllContainerByID() []string{
	cmd := exec.Command("docker","ps","-a","-q")
	out,err := cmd.CombinedOutput()
	
	if err!= nil{
		panic(err.Error())
		
	}
	str_out := strings.TrimSpace(string(out))
  
	var str []string = strings.Split(str_out,"\n")
	return str
}

func ListAllRunningByID() []string{
	cmd := exec.Command("docker","ps","-q")
	out,err := cmd.CombinedOutput()
	
	if err!= nil{
		fmt.Println(err.Error())
		return nil
	}
	str_out := strings.TrimSpace(string(out))
  
	var str []string = strings.Split(str_out,"\n")
	return str
}

func DeleteById(containerID string) string{
	cmd := exec.Command("docker","rm", containerID)
	out,err := cmd.CombinedOutput()
	str_out := string(out)
	if err!= nil{
		panic(err.Error())
		
	}
	return str_out
}

func StopAll(containers []string)  {
	for i,item := range  containers{
		cmd := exec.Command("docker","stop",item)
		out,err := cmd.CombinedOutput()
		str_out := string(out)
		if err!= nil{
			panic(err.Error())
		}
		fmt.Println("Parado Com Sucesso: ",i ,"->",str_out)
	}

}