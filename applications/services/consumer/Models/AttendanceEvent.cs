namespace LSExam.Models;
public record AttendanceEvent(string AttendanceCode, string StudentId, long CurrentUnixTime, double Latitude, double Longitude)
{
    public AttendanceEvent() : this("", "", 0, 0, 0) {}
}