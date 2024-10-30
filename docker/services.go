package docker
import "os/exec"



func FireFox() (string, string){
	cmd,err := exec.Command(
		"docker", "run", "-d", "-P",
        "--security-opt", "seccomp=unconfined",
        "-e", "TZ=Etc/UTC",
        "--shm-size=1gb",
        "--restart", "unless-stopped",
        "lscr.io/linuxserver/firefox:latest",
		).Output()

	if err != nil{
		panic(err.Error())
	}
	str_out := string(cmd)

	result := str_out[0:12]
	out,_ := exec.Command("docker","port",result).Output()

	port:= string(out)
	
	return port,result
}

func Minecraft() (string,string)  {
	cmd,err := exec.Command(
		"docker",
		 "run", 
		 "-d",
		 "-it",
		 "-e",
		 "EULA=TRUE",
		  "-P",
		  "-v", 
		  "mc-bedrock-data:/data", 
		  "itzg/minecraft-bedrock-server",

	).Output()
	if err != nil{
		panic(err.Error())
	}
	str_out := string(cmd)

	result := str_out[0:12]
	out,_ := exec.Command("docker","port",result).Output()

	port:= string(out)
	
	return port,result
}