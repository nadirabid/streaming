import { useEffect } from 'react';
import { Provider } from 'react-redux'
import "@fortawesome/fontawesome-svg-core/styles.css"; // import Font Awesome CSS
import { config } from "@fortawesome/fontawesome-svg-core";

import '../styles/globals.scss'
import { useStore } from '../store'
import { configureScreenSizeMediaQuery } from '../store/screenSize/actions'

config.autoAddCss = false;

export default function App({ Component, pageProps }) {
  const store = useStore(pageProps.initialReduxState)

  useEffect(() => {
    configureScreenSizeMediaQuery(store.dispatch)
  })

  return (
    <Provider store={store}>
      <Component {...pageProps} />
    </Provider>
  )
}
