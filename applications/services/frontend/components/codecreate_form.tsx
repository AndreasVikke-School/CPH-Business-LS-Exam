import { FormEvent } from "react"
import Router from 'next/router'

const CodeCreateForm = () => {
    return (
        <form onSubmit={checkin}>
            <h3>Create New Attendance Code</h3>
            <hr />
            <div className="row g-3 align-items-center">
                <div className="col-auto">
                    <label className="col-form-label">Minutes to live:</label>
                </div>
                <div className="col-auto">
                    <input type="number" id="attendance_code" className="form-control" placeholder="minutes must be a number" />
                </div>
                <div className="col-auto">
                    <button type="submit" className="btn btn-primary">Generate Code</button>
                </div>
            </div>
        </form>
    )
}
export default CodeCreateForm

const checkin = async (event: FormEvent) => {
    event.preventDefault()

    const res = await fetch('http://api:8080/api/attendance_code/' + event.target.attendance_code.value, {
        headers: {
            'Content-Type': 'application/json'
        },
        method: 'POST'
    })
    const code = await res.json()
    localStorage.setItem("code", JSON.stringify(code))
    Router.push('/code_show')
}