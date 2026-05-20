<script setup lang="ts">
    import { reactive, onMounted, onBeforeUnmount } from 'vue';
    import { useI18n } from "vue-i18n";

    import { NModal } from 'naive-ui';

    import { useSessionStore } from '../../../stores/session';
    import { appBus } from '../../../shared/composables/bus';
    import { TokenManager } from '../services/tokenManager';
    import LoginForm from './LoginForm.vue';

    const { t } = useI18n();
    const sessionStore = useSessionStore();

    let stopBusReauthListener: () => void;

    const reAuthEmitters = reactive<Array<string>>([]);

    const show = defineModel('show', { default: false });

    const onSuccessReauth = () => {
        show.value = false;
        appBus.emit({ type: "reauthValidNotify", payload: { to: reAuthEmitters } });
        reAuthEmitters.length = 0;
    };

    onMounted(async () => {
        stopBusReauthListener = appBus.on("reauthRequired", async (payload) => {
            reAuthEmitters.push(payload.emitter);
            try {
                const success = await TokenManager.refreshAccessToken(sessionStore);
                if (success) {
                    onSuccessReauth();
                } else {
                    show.value = true;
                }
            } catch (error: unknown) {
                console.error("An unhandled exception occurred during access token refresh", error);
                show.value = true;
            };
        });
    });

    onBeforeUnmount(() => {
        stopBusReauthListener();
    });
</script>

<template>
    <n-modal :title="t('modules.auth.components.reauthModal.title')" v-model:show="show" preset="card"
        class="doneo-reauth-modal">
        <LoginForm @success="onSuccessReauth" />
    </n-modal>
</template>

<style lang="css" scoped>
    .doneo-reauth-modal {
        width: 400px;
    }
</style>