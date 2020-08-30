import Head from 'next/head';
import { useState } from 'react';
import styles from '../styles/Home.module.scss';

export default function Home(): JSX.Element {
  return (
    <div className={styles.container}>
      <Head>
        <title>Curve</title>
      </Head>

      <main className={styles.main}>
        <section className={styles["top-section"]}>
          <div className={styles.title}>
            Curve
            <div className={styles['sub-title']}>
              探索未知的世界。<br />
              拓展您的现在和未来。
            </div>
          </div>

          <div className={styles['sign-up-form-container']}>
            <div>现在加入我们。</div>
            <div>已有账号？<a href="/login">现在登录。</a></div>
            <form className={styles['sign-up-form']}>
              <input id="email" type="text" placeholder="邮箱地址"></input>
              <input id="verification-code" type="text" placeholder="验证码"></input>
              <input id="password" type="password" placeholder="密码"></input>
              <input id="password-again" type="password" placeholder="确认密码"></input>
              <div>
                <input type="checkbox" id="read-agreement"></input>
                <label htmlFor="read-agreement">我已阅读并同意《Curve 用户协议》。</label>
              </div>
              <input type="submit" value="注册" />
            </form>
          </div>
        </section>
      </main>

      <footer className={styles.footer}>
        <p>Web 秃头小分队 - 2020 登陆项目</p>
      </footer>
    </div>
  );
}
