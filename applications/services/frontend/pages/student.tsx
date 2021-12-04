import type { NextPage } from 'next'
import Head from 'next/head'
import { getSession, useSession } from "next-auth/react";
import CheckInTable from '../components/checkins_table'
import CheckInForm from '../components/checkin_form'
import Menu from '../components/menu'
import styles from '../styles/Home.module.css'
import Link from 'next/link'
import React, { useCallback, useEffect, useState } from 'react';
import Home from '.';

const Student = () => {
    const [data, setData] = useState({})
    const { data: session, status } = useSession()
    const isUser = !!session?.user

    const fetchData = useCallback(async () => {
        const res = await fetch(`http://${process.env.NEXT_PUBLIC_API_IP}/api/checkins/student/${session?.user?.email}`);
        const data = await res.json();
        setData(data)
    }, [isUser, status])

    useEffect(() => {
        if (status === "loading") return
        fetchData()
    }, [isUser, status])

    if (isUser)
        return (
            <div className={styles.container}>
                <Head>
                    <title>AA - Student Page</title>
                    <link rel="icon" href="/favicon.ico" />
                </Head>

                <Menu />

                <main className={styles.main}>
                    <h1 className={styles.title}>
                        Welcome to <Link href="/student"><a>Student</a></Link> page
                    </h1>

                    <p className={styles.description}>
                        Logged in as{' '}
                        <code className={styles.code}>{session?.user?.name}</code>
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
    else
        return (
            <Home />
        )
}

export default Student
