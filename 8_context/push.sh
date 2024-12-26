# 所有的服务器列表
#servers=(  "n212" "n213" "n214" "n215" "n219"
#"n220" "n221" "n222" "n223" "n224" "n225" "n226" "n227" "n228" "n229"
#"n230" "n231" "n232" "n233" "n234" "n235" "n236")




#servers=("n206" "n207")
#servers=("n208" "n209" "n211" "n212")
#servers=("n213" "n214" "n215")
#servers=("n216" "n217" "n218" "n219")
#servers=("n220" "n221" "n222")
#servers=("n223" "n224" "n225" "n226")
#servers=("n227" "n228" "n229" "n230")
servers=("n231" "n232" "n234" "n235" "n236")

password="QZMPwdb13568963819"

for server in "${servers[@]}"
do
#    sshpass -p "$password" ssh -o StrictHostKeyChecking=no root@"$server" "sudo yum install atop"
    sshpass -p "$password" ssh -o StrictHostKeyChecking=no root@"$server" "sudo systemctl restart atop"
done

# 采集服务 安装系统级日志监控