package main

/*
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/socket.h>
int sock;
char* chars=NULL;
void* SendAndRece(char* msgSend){
	chars=NULL;
	char sendLine[300];
	strcpy(sendLine,msgSend);
	send(sock,msgSend,strlen(sendLine),0);
	char buffer[10240];
	int n=recv(sock,buffer,10240,0);
	buffer[n]='\0';
	printf("%s\n",buffer);
	//memset(buffer,0,sizeof(buffer));
	char* chars=buffer;
}
void *Client(){
	FILE *stream;
	char line[20];
	if( (stream = fopen( "ipaddress", "r" )) != NULL )
   {
	   if(fgets(line,20,stream)==NULL)
	   {
		printf("check ipaddress");
	   }
      fclose( stream );
   }
	sock=socket(AF_INET,SOCK_STREAM,0);
	if(sock<0){
		printf("Create Socket Error\n");
	}
	struct sockaddr_in serv_addr;
	memset(&serv_addr,0,sizeof(serv_addr));
	serv_addr.sin_family=AF_INET;
    serv_addr.sin_addr.s_addr=inet_addr(line);
    serv_addr.sin_port=htons(8000);
	if(connect(sock,(struct sockaddr*)&serv_addr,sizeof(serv_addr))<0){
        printf("Server Connect Error\n");
        return  0;
    }
}
void *Close(){
	close(sock);
}
*/
import "C"
import "os"
func GoSendAndRecv(msg string) string{
	cs := C.CString(msg)
	C.SendAndRece(cs)
	str :=C.GoString(C.chars)
	return str
}
var args string
func Init(){
	C.Client()
}
func GoCloseClient(){
	C.Close()
}
func main(){
	Init()
	args=os.Args[1]
 //       args:="CreateContainer-tam-tamp"
	GoSendAndRecv(args)
	GoCloseClient()

}

