import Head from 'next/head'
import styles from '../styles/Home.module.css'
import Image from 'next/image'

export default function Home() {
  let rows = []

  for (let i = 0; i < 20; i++) {
    rows.push(
      <div className={styles.row}>
        <Columns />
      </div>
    )
  }

  return (
    <div className={styles.container}>
      <Head>
        <title>Streaming</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <div className={styles.header}>
          <span>
            Streaming
          </span>

          Home

          Movies

          TV Shows

          Latest
        </div>

        <div className={styles.content}>
          {rows}
        </div>
      </main>
    </div>
  )
}



function Columns() {
  const images = [1,2,3,4,5,6,7].map(() =>
    <span className={styles.imageTest}>
      <Image
        className={styles.image}
        height="142"
        width="250"
        layout="intrinsic"
        src="http://localhost:1234/assets/content/summer_adrift/one/thumbnails/842x480.jpg"
      />
    </span>
  )

  return (
    <div className={styles.test}>
      {images}
    </div>
  )
}

