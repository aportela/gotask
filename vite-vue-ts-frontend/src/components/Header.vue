<script setup lang="ts">
    import { IconUserCircle, IconBrandGithub, IconDatabaseStar, IconHome, IconSitemap, IconBug, IconFileAnalytics, IconSettings } from '@tabler/icons-vue';
</script>

<template>
    <header class="navbar navbar-expand-md d-print-none">
        <div class="container-xl">
            <!-- BEGIN NAVBAR TOGGLER -->
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbar-menu"
                aria-controls="navbar-menu" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <!-- END NAVBAR TOGGLER -->
            <!-- BEGIN NAVBAR LOGO -->
            <div class="navbar-brand navbar-brand-autodark d-none-navbar-horizontal pe-0 pe-md-3">
                <IconDatabaseStar />
                GoTask
            </div>
            <!-- END NAVBAR LOGO -->
            <div class="navbar-nav flex-row order-md-last">
                <div class="nav-item d-none d-md-flex ">
                    <div class="btn-list">
                        <a href="https://github.com/aportela/gotask" class="nav-link px-0" target="_blank"
                            rel="noreferrer">
                            <IconBrandGithub />
                        </a>
                    </div>
                </div>
                <div class="d-none d-md-flex">
                    <div class="nav-item">
                        <a href="?theme=dark" class="nav-link px-0 hide-theme-dark" data-bs-toggle="tooltip"
                            data-bs-placement="bottom" aria-label="Enable dark mode"
                            data-bs-original-title="Enable dark mode">
                            <!-- Download SVG icon from http://tabler.io/icons/icon/moon -->
                            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                                fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round" class="icon icon-1">
                                <path
                                    d="M12 3c.132 0 .263 0 .393 0a7.5 7.5 0 0 0 7.92 12.446a9 9 0 1 1 -8.313 -12.454z">
                                </path>
                            </svg>
                        </a>
                        <a href="?theme=light" class="nav-link px-0 hide-theme-light" data-bs-toggle="tooltip"
                            data-bs-placement="bottom" aria-label="Enable light mode"
                            data-bs-original-title="Enable light mode">
                            <!-- Download SVG icon from http://tabler.io/icons/icon/sun -->
                            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                                fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                stroke-linejoin="round" class="icon icon-1">
                                <path d="M12 12m-4 0a4 4 0 1 0 8 0a4 4 0 1 0 -8 0"></path>
                                <path
                                    d="M3 12h1m8 -9v1m8 8h1m-9 8v1m-6.4 -15.4l.7 .7m12.1 -.7l-.7 .7m0 11.4l.7 .7m-12.1 -.7l-.7 .7">
                                </path>
                            </svg>
                        </a>
                    </div>
                </div>
                <div class="nav-item dropdown">
                    <a href="#" class="nav-link d-flex lh-1 p-0 px-2" data-bs-toggle="dropdown"
                        aria-label="Open user menu">
                        <IconUserCircle :size="24" />
                        <!--
                        <span class="avatar avatar-sm" style="background-image: url(./static/avatars/000m.jpg)"> </span>
                        -->
                        <div class="d-none d-xl-block ps-2">
                            <div>John Doe</div>
                            <div class="mt-1 small text-secondary">Tester</div>
                        </div>
                    </a>
                    <div class="dropdown-menu dropdown-menu-end dropdown-menu-arrow">
                        <a href="#" class="dropdown-item">Status</a>
                        <a href="./profile.html" class="dropdown-item">Profile</a>
                        <a href="#" class="dropdown-item">Feedback</a>
                        <div class="dropdown-divider"></div>
                        <a href="./settings.html" class="dropdown-item">Settings</a>
                        <a href="./sign-in.html" class="dropdown-item">Logout</a>
                    </div>
                </div>
            </div>
        </div>
    </header>
    <header class="navbar-expand-md">
        <div class="collapse navbar-collapse" id="navbar-menu">
            <div class="navbar">
                <div class="container-xl">
                    <div class="row flex-column flex-md-row flex-fill align-items-center">
                        <div class="col">
                            <!-- BEGIN NAVBAR MENU -->
                            <ul class="navbar-nav">
                                <li class="nav-item active">
                                    <a class="nav-link" href="./">
                                        <span class="nav-link-icon d-md-none d-lg-inline-block">
                                            <IconHome />
                                        </span>
                                        <span class="nav-link-title"> Home </span>
                                    </a>
                                </li>
                                <li class="nav-item">
                                    <a class="nav-link" href="./">
                                        <span class="nav-link-icon d-md-none d-lg-inline-block">
                                            <IconSitemap />
                                        </span>
                                        <span class="nav-link-title"> Projects </span>
                                    </a>
                                </li>
                                <li class="nav-item">
                                    <a class="nav-link" href="./">
                                        <span class="nav-link-icon d-md-none d-lg-inline-block">
                                            <IconBug />
                                        </span>
                                        <span class="nav-link-title"> Tasks </span>
                                    </a>
                                </li>
                                <li class="nav-item">
                                    <a class="nav-link" href="./">
                                        <span class="nav-link-icon d-md-none d-lg-inline-block">
                                            <IconFileAnalytics />
                                        </span>
                                        <span class="nav-link-title"> Reports </span>
                                    </a>
                                </li>
                                <li class="nav-item">
                                    <a class="nav-link" href="./">
                                        <span class="nav-link-icon d-md-none d-lg-inline-block">
                                            <IconSettings />
                                        </span>
                                        <span class="nav-link-title"> Settings </span>
                                    </a>
                                </li>
                            </ul>
                            <!-- END NAVBAR MENU -->
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </header>
</template>