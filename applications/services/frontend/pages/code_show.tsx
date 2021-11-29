import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import React, { useEffect, useState } from 'react'
import Countdown from 'react-countdown'
import styles from '../styles/Home.module.css'

const CodeShow: NextPage = () => {
    const [code, setCode] = useState({});

    useEffect(() => {
        setCode(JSON.parse(localStorage.getItem("code") || '{}'))
    }, []);

    const renderer = ({ minutes, seconds, completed }) => {
        if (completed) {
            return <span className="badge bg-danger">TIMES UP</span>;
        } else {
            return <span className="badge bg-success">{minutes}:{seconds}</span>;
        }
    };

    return (
        <div className={styles.container}>
            <Head>
                <title>Create Next App</title>
                <meta name="description" content="Generated by create next app" />
                <link rel="icon" href="/favicon.ico" />
            </Head>

            <main className={styles.main}>
                <h1 className={styles.title}>
                    <a href="/code_show">{code.code}</a>
                </h1>

                <p className={styles.description}>
                    {code.unix}
                    <br />
                    {Date.now() + (code.unix - Date.now())}
                    <br />
                    {code.unix}
                    <br />
                    {Date.now()}
                    <br />
                    Time left: <Countdown date={1638226962887} renderer={renderer} />
                </p>
            </main>
        </div>
    )
}

export default CodeShow