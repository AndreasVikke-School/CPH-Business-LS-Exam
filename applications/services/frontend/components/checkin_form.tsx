import { FormEvent } from "react"
import Router from 'next/router'

const CheckInForm = () => {
    return (
        <form onSubmit={checkin}>
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

const checkin = async (event: FormEvent) => {
    event.preventDefault()
}