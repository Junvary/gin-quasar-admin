<template>
    <div class="scroll-down-container">
        <div class="chevron"></div>
        <div class="chevron"></div>
        <div class="chevron"></div>
    </div>
</template>

<script setup></script>

<style lang="scss" scoped>
.scroll-down-container {
    position: relative;
    width: 30px;
    height: 24px;

    &:hover {
        cursor: pointer;
    }
}

.chevron {
    position: absolute;
    width: 30px;
    height: 6px;
    opacity: 0;
    transform: scale3d(0.5, 0.5, 0.5);
    animation: move 3s ease-out infinite;
}

.chevron:first-child {
    animation: move 3s ease-out 1s infinite;
}

.chevron:nth-child(2) {
    animation: move 3s ease-out 2s infinite;
}

.chevron:before,
.chevron:after {
    content: ' ';
    position: absolute;
    top: 0;
    height: 100%;
    width: 51%;
    background: #fff;
}

.chevron:before {
    left: 0;
    transform: skew(0deg, 30deg);
}

.chevron:after {
    right: 0;
    width: 50%;
    transform: skew(0deg, -30deg);
}

@keyframes move {
    25% {
        opacity: 1;
    }

    33% {
        opacity: 1;
        transform: translateY(30px);
    }

    67% {
        opacity: 1;
        transform: translateY(40px);
    }

    100% {
        opacity: 0;
        transform: translateY(55px) scale3d(0.5, 0.5, 0.5);
    }
}

@keyframes pulse {
    to {
        opacity: 1;
    }
}
</style>