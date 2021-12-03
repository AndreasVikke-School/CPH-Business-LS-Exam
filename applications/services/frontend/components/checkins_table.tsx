type tplotOptions = {
    [key: string]: string
}

const statuses: tplotOptions = {
    "1": "success",
    "2": "outOfTime",
    "3": "notFound",
    "4": "error"
}

const CheckInTable = ({ data }: { data: any }) => {
    return (
        <table className="table">
            <thead>
                <tr>
                    <th scope="col">Attendance Code</th>
                    <th scope="col">CheckIn Time</th>
                    <th scope="col">Status</th>
                </tr>
            </thead>
            <tbody>
                {data.map((d:any) => {
                    d.status = 2
                    var c = d.status == 1 ? "success" : d.status == 2 ? "warning" : "danger"
                    var date = new Date(d.checkinTime)
                    return (
                        <tr key={d.attendance_code}>
                            <th scope="row">{d.attendanceCode}</th>
                            <td>{date.toDateString()} - {date.getHours()}:{date.getMinutes()}</td>
                            <td><span className={"badge bg-" + c}>{statuses[d.status]}</span></td>
                        </tr>
                    )
                })}
            </tbody>
        </table>
    )
}
export default CheckInTable