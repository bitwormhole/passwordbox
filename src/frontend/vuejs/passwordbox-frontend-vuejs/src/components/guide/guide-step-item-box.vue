<script lang="js">
import { ElButton } from 'element-plus';



export default {

    computed: {
        current_mark() {
            if (this.step == this.current) {
                return "*";
            }
            return "";
        },
    },

    data() {
        const foo = ""
        return { foo }
    },

    methods: {

        gotoNextStep() {
            let value = this.current + 1;
            this.$emit('update:current', value);
        },

        gotoPrevStep() {
            let value = this.current - 1;
            this.$emit('update:current', value);
        },
    },

    props: {
        title: String,
        step: Number,
        current: Number, // current-step
    }

}

</script>

<template>

    <div class="step-item-box">
        <div class="step-item-title-bar">
            {{ current_mark }} 第 {{ step }} 步: {{ title }}
        </div>
        <div class="step-item-client-box">
            <slot> </slot>

            <div class="step-button-bar">
                <ElButton class="step-button-bar-button" type="default" @click="gotoPrevStep"> 上一步 </ElButton>
                <span class="step-button-bar-spring"> </span>
                <ElButton class="step-button-bar-button" type="primary" @click="gotoNextStep"> 下一步 </ElButton>
            </div>
        </div>
    </div>

</template>

<style lang="css">
.step-item-title-bar {

    background-color: rgb(152, 152, 197);
    margin: 5px;
    padding: 5px;
    color: rgb(235, 235, 235);
}

.step-item-client-box {
    padding: 10px;
}

.step-button-bar {
    display: flex;
    flex-direction: row;
    margin-top: 20px;
}

.step-button-bar-button {
    flex-grow: 0;
}

.step-button-bar-spring {
    flex-grow: 1;
}
</style>
