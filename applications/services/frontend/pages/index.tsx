import type { NextPage } from 'next'
import Head from 'next/head'
import React, { useEffect } from 'react'
import Menu from '../components/menu'
import styles from '../styles/Home.module.css'
import Link from 'next/link'
import LoginForm from '../components/login_form'

const Home: NextPage = () => {
  return (
    <div className={styles.container}>
      <Head>
        <title>Create Next App</title>
        <meta name="description" content="AP - Home" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <Menu />

      <main className={styles.main}>
        <h1 className={styles.title}>
          Welcome to <Link href="/"><a>Attendance App</a></Link>
        </h1>
        <p className={styles.description}>
          Use GitHub OAuth to authenticate
        </p>
          <LoginForm />
      </main>
    </div>
  )
}

export default Home
