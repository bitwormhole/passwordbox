<script>
import { ElButton, ElInput } from 'element-plus';


export default {

    data() {
        const email_address = ""
        return { email_address }
    }

}


</script>

<template>
    <div>

        <h2> 创建新账号 </h2>


        <ElInput v-model="email_address" placeholder="Email 地址"></ElInput>

        <ElButton type="primary"> 下一步 </ElButton>
    </div>
</template>
