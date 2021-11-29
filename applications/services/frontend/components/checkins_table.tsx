

const CheckInTable = ({ data }) => {
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
                {data.map((d) => {
                    var c = d.status == "success" ? "success" : d.status == "oot" ? "warning" : "danger"
                    var date = new Date(d.unix * 1000)
                    return (
                        <tr key={d.attendance_code}>
                            <th scope="row">{d.attendance_code}</th>
                            <td>{date.toDateString()} - {date.getHours()}:{date.getMinutes()}</td>
                            <td><span className={"badge bg-" + c}>{d.status}</span></td>
                        </tr>
                    )
                })}
            </tbody>
        </table>
    )
}
export default CheckInTable