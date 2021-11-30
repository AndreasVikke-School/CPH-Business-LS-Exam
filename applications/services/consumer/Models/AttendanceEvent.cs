namespace LSExam.Models;
public record AttendanceEvent(string AttendanceCode, string StudentId, long CurrentUnixTime)
{
    public AttendanceEvent() : this("", "", 0) {}
}