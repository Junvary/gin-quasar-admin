import { ref } from 'vue';

export default function useVideoContainer() {
    const videoCanPlay = ref(false);
    const fixStyle = ref('')
    const ramdomSrc = `video/video-${Math.floor(Math.random() * 7) + 1}.mp4`
    const canplay = () => {
        videoCanPlay.value = true
    }
    //屏幕自适应
    const resizeWindow = () => {
        window.onresize = () => {
            const windowWidth = document.body.clientWidth
            const windowHeight = document.body.clientHeight
            const windowAspectRatio = windowHeight / windowWidth
            let videoWidth
            let videoHeight
            if (windowAspectRatio < 0.5625) {
                videoWidth = windowWidth
                videoHeight = videoWidth * 0.5625
                fixStyle.value = {
                    height: windowWidth * 0.5625 + 'px',
                    width: windowWidth + 'px',
                    'margin-bottom': (windowHeight - videoHeight) / 2 + 'px',
                    'margin-left': 'initial'
                }
            } else {
                videoHeight = windowHeight
                videoWidth = videoHeight / 0.5625
                fixStyle.value = {
                    height: windowHeight + 'px',
                    width: windowHeight / 0.5625 + 'px',
                    'margin-left': (windowWidth - videoWidth) / 2 + 'px',
                    'margin-bottom': 'initial'
                }
            }
        }
        window.onresize()
    }
    return {
        resizeWindow,
        ramdomSrc,
        videoCanPlay,
        canplay,
        fixStyle
    }
}