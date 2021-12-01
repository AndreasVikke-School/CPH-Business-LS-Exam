import { FormEvent, useState } from "react"
import Router from 'next/router'

const LoginForm = () => {
    const [userId, setUserId] = useState("")

    const teacher = async (event: FormEvent) => {
        localStorage.setItem("teacherId", userId)
        Router.push("/teacher")
    }

    const student = async (event: FormEvent) => {
        localStorage.setItem("studentId", userId)
        Router.push("/student")
    }
    
    return (
        <form>
            <h3>Login</h3>
            <hr />
            <div className="row g-3 align-items-center">
                <div className="col-auto">
                    <label className="col-form-label">User Id:</label>
                </div>
                <div className="col-auto">
                    <input onChange={(e) => setUserId(e.target.value)} required type="text" id="attendance_code" className="form-control" placeholder="user id" />
                </div>
            </div>
            <br />
            <div className="row g-3 align-items-center">
                <div className="col-auto">
                    <button type="submit" onClick={teacher} className="btn btn-primary">Login as Teacher</button>
                </div>
                <div className="col-auto">
                    <button type="submit" onClick={student} className="btn btn-primary">Login as Student</button>
                </div>
            </div>
        </form>
    )
}
export default LoginForm