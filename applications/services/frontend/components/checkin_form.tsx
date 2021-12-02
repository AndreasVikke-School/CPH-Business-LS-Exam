import React, { FormEvent, useState } from "react"
import ErrorBadge from "./error_badge"
import { useSession } from "next-auth/react";

const CheckInForm = () => {
    const { data: session } = useSession()
    const [status, setStatus] = useState(0)

    const checkin = async (event: FormEvent) => {
        event.preventDefault()
        var lat = 0, long = 0

        if ("geolocation" in navigator) {
            navigator.geolocation.getCurrentPosition(function (position) {
                lat = position.coords.latitude;
                long = position.coords.longitude;
            });
        } else {
            console.log("GPS Not Available");
        }

        const res = await fetch(`http://${process.env.NEXT_PUBLIC_API_IP}/api/checkin/`, {
            body: JSON.stringify({
                attendanceCode: (event.target as any).attendance_code.value,
                studentId: session?.user?.email,
                lat: lat,
                long: long
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