# pbox:UserBank

UserBank 表示用户的密码存储空间(文件夹).

它们的文件夹结构为:

1. 根用户:

        ~/.passwordbox/root/bank/
        ~/.passwordbox/root/bank/refs
        ~/.passwordbox/root/bank/blocks
        ~/.passwordbox/root/bank/properties

2. 常规用户:

        ~/.passwordbox/home/{user_name}/bank/
        ~/.passwordbox/home/{user_name}/bank/refs
        ~/.passwordbox/home/{user_name}/bank/blocks
        ~/.passwordbox/home/{user_name}/bank/properties
