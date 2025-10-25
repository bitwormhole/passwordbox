<script lang="js">
import { ElButton, ElForm, ElFormItem, ElInput } from 'element-plus';


import { useAxiosStore } from '@/stores/lib_axios'

const theAxiosStore = useAxiosStore()


export default {

    data() {

        const item = {
            email: 'demo1@example.com',
            domain: '',
            username: '',
            scene: 'default',
            revision: 0,
            length: 11,
            charset: 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz~!@#$%^&*()_+{}|:<>?-=[];,.',
            verification: ''
        }

        return { item }
    },


    methods: {
        handleClickGenerate() {

            let item = this.item;

            let method = 'POST'
            let url = '/api/v1/passwords/do/fast-gen'
            let data = { passwords: [item] }
            let config = { method, url, data }
            let p = theAxiosStore.execute(config);
        },
    }

}

</script>

<template>
    <div>
        <ElForm label-width="120">
            <ElFormItem label="Email" required>
                <ElInput v-model="item.email" placeholder="请输入"></ElInput>
            </ElFormItem>
            <ElFormItem label="Domain" required>
                <ElInput v-model="item.domain" placeholder="请输入"></ElInput>
            </ElFormItem>
            <ElFormItem label="UserName" required>
                <ElInput v-model="item.username" placeholder="请输入"></ElInput>
            </ElFormItem>
            <ElFormItem label="Scene">
                <ElInput v-model="item.scene" placeholder="请输入"></ElInput>
            </ElFormItem>

            <ElFormItem label="Revision">
                <ElInputNumber v-model="item.revision" :min="0" :max="999999"></ElInputNumber>
            </ElFormItem>

            <ElFormItem label="Length">
                <ElInputNumber v-model="item.length" :min="3" :max="33"></ElInputNumber>
            </ElFormItem>
            <ElFormItem label="CharSet">
                <ElInput v-model="item.charset"></ElInput>
            </ElFormItem>

            <ElFormItem label="Verification" required>
                <ElInput v-model="item.verification" type="password" placeholder="请输入验证码"></ElInput>
            </ElFormItem>

            <ElFormItem label="">
                <div>
                    <ElButton type="primary" @click="handleClickGenerate"> 生成 </ElButton>
                </div>
            </ElFormItem>

        </ElForm>



    </div>
</template>
