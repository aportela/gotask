<script setup lang="ts">
    import { NAlert } from 'naive-ui'

    type AlertType = "default" | "info" | "warning" | "error" | "success" | undefined;

    interface RemoteAPIAlertProps {
        title: string,
        type: AlertType,
        message: string
    };

    withDefaults(defineProps<RemoteAPIAlertProps>(), {
        title: "",
        type: undefined,
        message: "",
    });
</script>

<template>
    <n-alert :title="title" :type="type">
        <span v-html="message.split('\n').map(line => `${line}<br />`).join('')"></span>
    </n-alert>
</template>

<style lang="css" scoped></style>