import Head from 'next/head';
import styles from '../styles/Home.module.scss';

export default function Home(): JSX.Element {
  return (
    <div className={styles.container}>
      <Head>
        <title>Curve</title>
        <link rel="icon" href="/icon.png" />
      </Head>

      <main className={styles.main}>
        <p>Hello, world!</p>
      </main>

      <footer className={styles.footer}>
        <p>Footer of the page</p>
      </footer>
    </div>
  );
}
