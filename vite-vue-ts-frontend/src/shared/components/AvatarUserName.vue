<script setup lang="ts">
    import { NAvatar } from 'naive-ui';
    import type { UserBase } from '../../modules/users/models/user';

    interface AvatarUserNameProps {
        user: UserBase;
        avatarSize?: number;
    };

    const props = withDefaults(defineProps<AvatarUserNameProps>(), {
        avatarSize: 32,
    });

    const avatarURL = `/api/avatars/${props.avatarSize}/user/${props.user.id}`;
</script>

<template>
    <div class="doneo-flex-center-align" style="gap: 8px;">
        <n-avatar :src="avatarURL" class="doneo-avatar-username" :size="props.avatarSize" />
        {{ user.name }}
    </div>
</template>

<style lang="css" scoped>
    .doneo-avatar-username {
        margin-right: 4px;
    }
</style>