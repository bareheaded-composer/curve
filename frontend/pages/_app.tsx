import '../styles/globals.scss';
import type { AppProps } from 'next/app';

function MyApp({ Component, pageProps }: AppProps): JSX.Element {
  /* eslint-disable react/jsx-props-no-spreading */
  return <Component {...pageProps} />;
}

export default MyApp;
