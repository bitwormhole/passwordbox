<script lang="js">

import { ElInput } from 'element-plus';
import MyStepListItem from './../guide-step-item-box.vue'
import MyStepListView from './../guide-step-list-view.vue'

export default {

    components: { MyStepListItem, MyStepListView },

    data() {
        const currentStep = 2;
        const email_address = '';
        return { currentStep, email_address }
    },


    methods: {

    }

}

</script>

<template>

    <div>

        <MyStepListView>

            <MyStepListItem v-model:current="currentStep" :step="1" title="填写 Email 地址">
                <ElInput v-model="email_address" type="" placeholder="请输入你的 Email 地址"></ElInput>
            </MyStepListItem>

            <MyStepListItem v-model:current="currentStep" :step="2" title="选择 Hub 服务商"> BBB </MyStepListItem>

            <MyStepListItem v-model:current="currentStep" :step="3" title="通过 Email 激活新账号"> BBB </MyStepListItem>

            <MyStepListItem v-model:current="currentStep" :step="4" title="绑定微信小程序"> CCC </MyStepListItem>

            <MyStepListItem v-model:current="currentStep" :step="5" title="完成"> CCC </MyStepListItem>

        </MyStepListView>

    </div>

</template>
