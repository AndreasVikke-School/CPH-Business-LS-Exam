import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import React from 'react'
import CheckInTable from '../components/checkins_table'
import CheckInForm from '../components/checkin_form'
import Menu from '../components/menu'
import styles from '../styles/Home.module.css'

const data = [
    {
        attendance_code: "6453785",
        unix: "1638187500",
        status: "success"
    },
    {
        attendance_code: "7564873",
        unix: "1638183900",
        status: "success"
    },
    {
        attendance_code: "1364859",
        unix: "1638180300",
        status: "oot"
    },
    {
        attendance_code: "2341562",
        unix: "1638176700",
        status: "error"
    },
    {
        attendance_code: "7564859",
        unix: "	1638173100",
        status: "success"
    }
]

const Student: NextPage = () => {
    return (
        <div className={styles.container}>
            <Head>
                <title>AA - Student Page</title>
                <link rel="icon" href="/favicon.ico" />
            </Head>

            <Menu />

            <main className={styles.main}>
                <h1 className={styles.title}>
                    Welcome to <a href="/student">Student</a> page
                </h1>

                <p className={styles.description}>
                    Logged in as{' '}
                    <code className={styles.code}>cph-av105</code>
                </p>

                <div className={styles.table}>
                    <CheckInForm />
                </div>

                <div className={styles.table}>
                    <CheckInTable data={data} />
                </div>
            </main>
        </div>
    )
}

export default Student
