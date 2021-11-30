import React, { FormEvent, useState } from "react"
import Router from 'next/router'
import ErrorBadge from "./error_badge"

const CheckInForm = () => {
    const [status, setStatus] = useState(0)

    const checkin = async (event: FormEvent) => {
        event.preventDefault()

        const res = await fetch(`http://${process.env.NEXT_PUBLIC_API_IP}/api/checkin/`, {
            body: JSON.stringify({
                attendanceCode: (event.target as any).attendance_code.value,
                studentId: "cph-av105"
            }),
            headers: {
                'Content-Type': 'application/json'
            },
            method: 'POST'
        })
        setStatus(res.status)
    }

    return (
        <form onSubmit={checkin}>
            <div className="row align-items-center mb-2"><ErrorBadge status={status} /></div>
            <div className="row g-3 align-items-center">
                <div className="col-auto">
                    <label className="col-form-label">Attendance Code:</label>
                </div>
                <div className="col-auto">
                    <input type="number" id="attendance_code" className="form-control" placeholder="Code must be 7 characters long." />
                </div>
                <div className="col-auto">
                    <button type="submit" className="btn btn-primary">CheckIn</button>
                </div>
            </div>
        </form>
    )
}
export default CheckInForm