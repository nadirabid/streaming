import Head from 'next/head'
import Image from 'next/image'
import { useState } from 'react'
import { connect } from 'react-redux'
import ChevronLeft from '@material-ui/icons/ChevronLeft';
import ChevronRight from '@material-ui/icons/ChevronRight';

import styles from '../styles/Home.module.scss'
import { screenSizes } from '../store/screenSize/actions'

export default function Home() {
  let rows = []

  for (let i = 0; i < 20; i++) {
    rows.push(
      <div className={styles.row}>
        <ConnectedContentRow />
      </div>
    )
  }

  return (
    <div className={styles.home}>
      <Head>
        <title>Qissa</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <div className={styles.navigationContainer}>
          <div className={styles.navigation}>
            <span>Qissa</span>
            <span>Movies</span>
            <span>TV Shows</span>
          </div>
        </div>

        <HeroContent />

        <div >
          {rows}
        </div>
      </main>
    </div>
  )
}

function HeroContent() {
  return (
    <div className={styles.heroContainer}>
      <img
        src="http://localhost:1234/assets/content/summer_adrift/one/thumbnails/842x480.jpg"
        className={styles.heroImage}
      />
      <div className={styles.heroText}>
        This is some text
      </div>
    </div>
  )
}

function ContentRow({ imageWidth, imageHeight }) {
  const [isShown, setIsShown] = useState(false)

  const images = [1,2,3,4,5,6].map((i) =>
    <Image
      key={i}
      className={styles.image}
      height={imageHeight}
      width={imageWidth}
      quality={100}
      layout="intrinsic"
      src="http://localhost:1234/assets/content/summer_adrift/one/thumbnails/842x480.jpg"
    />
  )

  return (
    <div
      className={styles.contentRowContainer}
      onMouseEnter={() => setIsShown(true)}
      onMouseLeave={() => setIsShown(false)}
    >
      {isShown && (<ChevronLeft className={styles.icon} />)}
      <div className={styles.rowsContainer}>{images}</div>
      {isShown && (<ChevronRight className={styles.icon} />)}
    </div>
  )
}

const ConnectedContentRow = connect(({ screenSize }) => {
  const sizeValue = screenSizes[screenSize];

  const numOfItems = 6
  const imageWidth = (sizeValue - (numOfItems - 1) * 8) / numOfItems 
  const imageHeight = imageWidth * 0.57

  return { imageWidth, imageHeight }
})(ContentRow)
