namespace LSExam.Services;

public class PostgresConsumer
{
    private readonly ILogger<PostgresConsumer> _logger;
    private readonly CheckInService.CheckInServiceClient _checkInClient;

    public PostgresConsumer(ILogger<PostgresConsumer> logger, CheckInService.CheckInServiceClient checkInClient)
    {
        _logger = logger;
        _checkInClient = checkInClient;
    }

    public async Task SaveAttendanceEvent(AttendanceEvent attendanceEvent, CodeValidity codeValidity)
    {
        CheckIn checkIn = new() { AttendanceCode = attendanceEvent.AttendanceCode, StudentId = attendanceEvent.StudentId, Status = codeValidity.ConvertToGrpcValidity(), CheckinTime = attendanceEvent.CurrentUnixTime };

        _ = await _checkInClient.InsertCheckInAsync(checkIn);
        _logger.LogInformation($"Attendance event was saved to db: {checkIn}");
    }
}
public static class Helper
{
    public static validity ConvertToGrpcValidity(this CodeValidity codeValidity)
        => codeValidity switch
        {
            CodeValidity.Success => validity.Success,
            CodeValidity.OutOfTime => validity.OutOfTime,
            CodeValidity.NotFound => validity.NotFound,
            CodeValidity.OutOfRange => validity.OutOfRange,
            CodeValidity.Error => validity.Error,
            _ => throw new NotImplementedException()
        };
}