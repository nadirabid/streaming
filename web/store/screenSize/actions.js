const xs = 0
const sm = 600
const md = 960
const lg = 1280
const xl = 1920

export const screenSizes = {
    xs,
    sm,
    md,
    lg,
    xl,
}

// this should be called AFTER redux store has been configured
export function configureScreenSizeMediaQuery(dispatch) {
    const xsQl = window.matchMedia(`(min-width:${xs}px) and (max-width:${sm}px)`)
    xsQl.addEventListener('change', () => dispatch(action('xs')))
    if (xsQl.matches) {
        dispatch(action('xs'))
    }

    const smQl = window.matchMedia(`(min-width:${sm}px) and (max-width:${md}px)`)
    smQl.addEventListener('change', () => dispatch(action('sm')))
    if (smQl.matches) {
        dispatch(action('sm'))
    }

    const mdQl = window.matchMedia(`(min-width:${md}px) and (max-width:${lg}px)`)
    mdQl.addEventListener('change', () => dispatch(action('md')))
    if (mdQl.matches) {
        dispatch(action('md'))
    }

    const lgQl = window.matchMedia(`(min-width:${lg}px) and (max-width:${xl}px)`)
    lgQl.addEventListener('change', () => dispatch(action('lg')))
    if (lgQl.matches) {
        dispatch(action('lg'))
    }

    const xlQl = window.matchMedia(`(min-width:${xl}px)`)
    xlQl.addEventListener('change', () => dispatch(action('xl')))
    if (xlQl.matches) {
        dispatch(action('xl'))
    }
}

export function action(value) {
    return {
        type: 'SCREEN_SIZE',
        screenSize: value,
    }
}
