<script setup lang="ts">
    import { useRouter } from 'vue-router'

    import { NResult, NSpace, NButton } from 'naive-ui'

    const router = useRouter()

    const goHome = () => {
        router.push('/')
    }

    const goBack = () => {
        router.back()
    }
</script>

<template>
    <div class="not-found">
        <n-result status="404" title="404 Not Found" description="We can't find what you're looking for." size="huge">
            <template #footer>
                <n-space justify="space-around">
                    <n-button type="primary" @click="goHome">
                        Go home
                    </n-button>
                    <n-button tertiary @click="goBack">
                        Go back
                    </n-button>
                </n-space>
            </template>
        </n-result>
    </div>
</template>

<style lang="css" scoped>
    .not-found {
        height: 100vh;
        display: flex;
        align-items: center;
        justify-content: center;
        padding: 24px;
    }
</style>