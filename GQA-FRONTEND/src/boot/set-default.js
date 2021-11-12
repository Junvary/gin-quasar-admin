import { Notify } from 'quasar'

Notify.setDefaults({
    progress: true,
    position: "top",
    classes: 'glossy',
    actions: [
        { icon: "close", color: 'white', handler: () => { /* ... */ } }
    ],
    // timeout: 3000
})