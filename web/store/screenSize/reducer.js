export default function reducer(state = 'lg', action) {
    switch (action.type) {
        case 'SCREEN_SIZE':
            return action.screenSize
        default:
            return state
    }
}
