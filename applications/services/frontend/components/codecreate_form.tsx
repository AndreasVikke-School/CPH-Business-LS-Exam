import { FormEvent } from "react"
import Router from 'next/router'
import { useSession } from "next-auth/react";

const CodeCreateForm = () => {
    const { data: session } = useSession()

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

        const res = await fetch(`http://${process.env.NEXT_PUBLIC_API_IP}/api/attendance_code/`, {
            body: JSON.stringify({
                minutesToLive: (event.target as any).minutesToLive.value,
                lat: lat,
                long: long
            }),
            headers: {
                'Content-Type': 'application/json'
            },
            method: 'POST'
        })
        const code = await res.json()
        localStorage.setItem("code", JSON.stringify(code))
        Router.push('/code_show')
    }

    return (
        <form onSubmit={checkin}>
            <h3>Create New Attendance Code</h3>
            <hr />
            <div className="row g-3 align-items-center">
                <div className="col-auto">
                    <label className="col-form-label">Minutes to live:</label>
                </div>
                <div className="col-auto">
                    <input type="number" id="minutesToLive" className="form-control" placeholder="minutes must be a number" />
                </div>
                <div className="col-auto">
                    <button type="submit" className="btn btn-primary">Generate Code</button>
                </div>
            </div>
        </form>
    )
}
export default CodeCreateForm