#include <stdio.h>
#include <stdlib.h>
#include <malloc.h>
#include <string.h>

#define string char*
#define int16 short
#define int32 int
#define int64 long
#define False 0
#define True 1
#define Error -1
#define bool int
#define Error_s "-1"

//#include "include.h"

/* List_string
    成员变量                                                                  说明
    int16 size                                                          表示list的当前有效长度，是长度，是data的index+1（在有1个数据的时候是1）
    int16 max_size                                                      表示当前申请的内存大小
    string[] data                                                       存储数据的地方
    string flag_data                                                    函数返回时携带的其他传递信息，如错误信息等

    成员函数:
    返回值 名字                         args                                说明                                        用例
    void  append            List_string *self, string x                 向list末尾添加变量                       list.append(&list,"in");
    bool  insert            List_string *self, string x ,int pos        向任意index插入变量，pos为其位置         list.insert(&list,"in",3);
    bool  del               List_string *self, int pos                  删除指定index的变量                      list.del(&list,3);
    List_string  copy       List_string *self                           独立复制一个相同的list                   List_string res = list.copy(&list);
    string  get             List_string *self, int pos                  从list中取出第pos位元素                  list.get(&list,3);
    bool  change            List_string *self, string x,int pos         修改pos位元素的值                        list.change(&list,"changed",3);
    int  find               List_string *self, string x ,int start_pos  在输入list中的指定位置开始查找指定元素，返回pos     list.find(&list,"in",0);
    void  connect           List_string *self, List_string *src         连接两个list，src是需要合并进self的list   list.connect(&list,&res);
    string list_to_string                                               list->string                             string x = list_to_string(&list);
    List_string string_to_list                                          string->list                              List_string res = string_to_list(x);
    void write_list_value        List_string *list                           list按行写成文件                          write_list_value(&list)
    

    声明方法：
    List_string new_list = new_list_string();   即可初始化一个list

*/

typedef struct List_String List_string;
void append_s(List_string *tmp,string x);
bool insert_s(List_string *tmp,string x,int16 pos);
bool del_s(List_string *tmp,int pos);
//List_string copy_s(List_string *tmp);  //这个方法可能会导致严重的内存溢出
string get_s(List_string *tmp,int pos);
bool change_s(List_string *tmp,string x,int pos);
int find_index_s(List_string *tmp,string x,int pos);
void connect_list(List_string *base,List_string *src);

typedef struct List_String{
    int16 size;
    int16 max_size;
    char** data;
    string flag_data;
    void(*append)(List_string *tmp,string x);
    bool(*insert)(List_string *tmp,string x,int16 pos);
    bool(*delete)(List_string *tmp,int pos);
    //List_string(*copy)(List_string *tmp);
    string(*get)(List_string *tmp,int pos);
    bool(*change)(List_string *tmp,string x,int pos);
    int (*find)(List_string *tmp,string x,int pos);
    void (*connect)(List_string *base,List_string *src);
}List_string;

List_string new_list_string(){
    List_string tmp;
    tmp.max_size = 50;
    tmp.size = 0;
    char** data_string;
    data_string = (string*)malloc(sizeof(string)*(tmp.max_size));
    tmp.flag_data = "None";
    tmp.data = data_string;
    tmp.append = *append_s;
    tmp.insert = *insert_s;
    tmp.delete = *del_s;
    //tmp.copy = *copy_s;
    tmp.get = *get_s;
    tmp.change = *change_s;
    tmp.find = *find_index_s;
    tmp.connect = *connect_list;
    return tmp;
}

void append_s(List_string *tmp,string x){

    int i = tmp->size;
    if(tmp->max_size == tmp->size-1){
        tmp->max_size += 20;
        string* data_string;
        data_string = (string*)malloc(sizeof(string)*(tmp->max_size+1));

        for(i=0;i<tmp->size;i++){
            data_string[i] = tmp->data[i];
        }
        free(tmp->data);
        tmp->data = data_string;

    }

    string data_tmp = (string)malloc(strlen(x)+1);
    strcpy(data_tmp,x);
    tmp->data[tmp->size] = data_tmp;
    tmp->size += 1;
}

bool insert_s(List_string *tmp,string x,int16 pos){
    if (tmp->size+1 < pos || pos < 0){
        return Error;
    }
    int i = 0;
    if(tmp->max_size <= tmp->size-2)
        tmp->max_size += 20;
    string* data_string;
    data_string = (string*)malloc(sizeof(string)*(tmp->max_size));
    for (i=0;i<pos;i++){
        data_string[i] = tmp->data[i];
    }
    data_string[pos] = x;
    tmp->size += 1;
    for(i=pos+1;i<tmp->size;i++){
        data_string[i] = tmp->data[i-1];
    }
    free(tmp->data);
    tmp->data = data_string;
    return True;

}

bool del_s(List_string *tmp,int pos){
    if (tmp->size < pos || pos < 0){
        return False;
    }
    string* data_string;
    data_string = (string*)malloc(sizeof(string)*(tmp->max_size));
    int i=0;
    for(i=0;i<pos;i++){
        data_string[i] = tmp->data[i];
    }
    for(i=pos+1;i<tmp->size;i++){
        data_string[i-1] = tmp->data[i];
    }
    free(tmp->data);
    tmp->data = data_string;
    tmp->size -= 1;
    return True;
}

/*
 * 这个方法可能会导致严重的内存溢出,暂时停用
 */
List_string copy_s(List_string *tmp){
    List_string new_list = new_list_string();
    free(new_list.data);
    string* data_string;
    new_list.max_size = tmp->max_size;
    data_string = (string*)malloc(sizeof(string)*(tmp->max_size));
    new_list.data = data_string;
    int i;
    for(i=0;i<tmp->size;i++){
        string x;
        x = (string)malloc(strlen(tmp->data[i]));
        strcpy(x,tmp->data[i]);
        new_list.append(&new_list,x);
        free(x);
    }
    return new_list;
}

bool change_s(List_string *tmp,string x,int pos){
    if (tmp->size < pos || pos < 0){
        return Error;
    }
    string* data_string;
    data_string = (string*)malloc(sizeof(string)*(tmp->max_size));
    int i=0;
    for(i=0;i<pos;i++){
        data_string[i] = tmp->data[i];
    }
    data_string[pos] = x;
    for(i=pos;i<tmp->size;i++){
        data_string[i-1] = tmp->data[i];
    }
    free(tmp->data);
    tmp->data = data_string;
    tmp->size -= 1;
    return True;
}

string get_s(List_string *tmp,int pos){
    if (tmp->size < pos || pos < 0){
        return Error_s;
    }

    return tmp->data[pos];
}

int find_index_s(List_string *tmp,string x,int start_pos){
    if(tmp->size == 0)
        return -3;
    if(start_pos > tmp->size || start_pos <0)
        return -2;
    int i;
    for(i=start_pos;i<tmp->size;i++){
        if(strcmp(tmp->data[i],x) == 0)
            return i;
    }
    return -1;
}

/*
 * 将src添加到base后面
 */
void connect_list(List_string *base,List_string *src){
    int i;
    for(i=0;i<src->size;i++) {
        base->append(base, src->get(src, i));
    }
}

void free_list(List_string *list){
    int i;
    for(i=0;i<list->size;i++){
        free(list->data[i]);
    }
    free(list->data);
    free(list);
}

/*
 * list转换成string
 * 已通过测试
 */
string list_to_string(List_string *list){
    int size_string=12;
    int i;
    char split_flag1[8] = "<split>";
    char split_flag2[9] = "</split>";
    char start_flag[8] = "<start>";
    char size_flag1[8] = "<size>";
    char size_flag2[9] = "</size>";
    char end_flag[9] = "</start>";
    for(i=0;i<list->size;i++){

        size_string = size_string + strlen(split_flag1)+ strlen(split_flag2) ;
        size_string = size_string + strlen(list->data[i]) ;
        size_string = size_string +9+9;
    }
    string res;
    res = (string)calloc(sizeof(char),size_string);

    strcpy(res,start_flag);
    for(i=0;i<list->size;i++){
        strcat(res,split_flag1);
        strcat(res,size_flag1);
        char x[10]={0};
        sprintf(x,"%d",(int)(strlen(list->data[i])));
        strcat(res,x);
        strcat(res,size_flag2);
        strcat(res,list->data[i]);


        strcat(res,split_flag2);

    }
    strcat(res,end_flag);
    return res;
}

/*
 * string转换成list
 * 已通过测试
 */
List_string string_to_list(string src){
    List_string res = new_list_string();
    int i;
    char split_flag1[8] = "<split>";
    char split_flag2[9] = "</split>";
    char start_flag[8] = "<start>";
    char size_flag1[8] = "<size>";
    char size_flag2[9] = "</size>";
    char end_flag[9] = "</start>";
    char f[20]={0};
    strncpy(f,src,7);
    if(strcmp(f,start_flag) == 0){
        src += 7; // 应该是7
    }
    int size=0;
    for(i=0;i<strlen(src);i++){
        strncpy(f,src,13);
        //printf("%s  \n",f);
        if(strcmp(f,"<split><size>") == 0){
            src += 13;
            int j;
            char tmp[6] = {0};
            char flag[8] = {0};
            for(j=0;j<strlen(src);j++){

                strncpy(flag,src,7);
                //printf("flag:%s  ",flag);
                if(strcmp(flag,"</size>") == 0){
                    size = atoi(tmp);
                    src+=7;
                    break;
                }
                char t[2]={0};
                strncpy(t,src,1);
                strcat(tmp,t);
                src+=1;
            }
        }else if(strcmp(f,"</split><spli") == 0){
            src += 21;
            int j;
            char tmp[6] = {0};
            char flag[8] = {0};
            for(j=0;j<strlen(src);j++){

                strncpy(flag,src,7);
                //printf("flag:%s  ",flag);
                if(strcmp(flag,"</size>") == 0){
                    size = atoi(tmp);
                    src+=7;
                    break;
                }
                char t[2]={0};
                strncpy(t,src,1);
                strcat(tmp,t);
                src+=1;
            }

        }
        //printf("size:%d  ",size);
        string data = (string)calloc(sizeof(char),size+1);
        int m;
        for(m=i;m<strlen(src)+200;m++){
            char flag[9] = {0};
            strncpy(flag,src,8);
            //printf("flag2:%s  ",flag);
            if(strcmp(flag,"</split>") == 0){
                //printf("data:%s\n",data);
                res.append(&res,data);
                src+=8;
                break;
            }
            char t[2]={0};
            strncpy(t,src,1);
            strcat(data,t);
            src+=1;
        }
        free(data);
        char e[9]={0};
        strncpy(e,src,8);
        //printf("src: %s\n",src);
        if(strcmp(e,"</start>") == 0)
            break;
    }

    return res;
}


void write_list_value(List_string *list){
    FILE *file = fopen("../temp.tmp","w");
    int i;
    for(i=0;i<list->size;i++){
        string data = calloc(sizeof(char),strlen(list->data[i])+13);
        strcat(data,list->data[i]);
        strcat(data,"\n");
        fputs(data,file);
        free(data);
    }
    fclose(file);
}
void print_list(List_string *list){
    int i;
    for(i=0;i<list->size;i++){
        printf("%s\n",list->data[i]);
    }
}
