<script setup lang="ts">
    import { default as TopHeader } from './TopHeader.vue';
    import { default as TopMenu } from './TopMenu.vue';
    import Header from '../components/Header.vue';
</script>

<template>
    <TopHeader />
    <TopMenu />
    <Header />
    <div class="page-wrapper">
        <div class="page-body">
            <div class="container-xl">
                <router-view />
            </div>
        </div>
    </div>
</template>