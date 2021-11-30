const ErrorBadge = ({ status } : { status: number}) => {
    if(status == 0)
        return null
    else if (status == 200)
        return <span className="badge bg-success">Attendance code has been queued successfully</span>
    else
        return <span className="badge bg-danger">Something went wrong, please try again</span>
}
export default ErrorBadge