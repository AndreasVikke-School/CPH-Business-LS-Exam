using Google.Protobuf.WellKnownTypes;
using LSExam.Configs;
using LSExam.Models;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Options;
using Protos;

namespace LSExam.Services;

public class RedisConsumer
{
    private readonly ILogger<RedisConsumer> _logger;
    private readonly AttendanceCodeProto.AttendanceCodeProtoClient _redisClient;

    public RedisConsumer(ILogger<RedisConsumer> logger, IOptions<RedisSettings> options, AttendanceCodeProto.AttendanceCodeProtoClient redisClient)
    {
        _logger = logger;
        _redisClient = redisClient;
    }

    public CodeValidity CheckCodeValidity(AttendanceEvent attendanceEvent)
    {
        long attandanceCode = long.Parse(attendanceEvent.AttendanceCode);
        Int64Value request = new() { Value = attandanceCode };
        try
        {
            AttendanceCode code = _redisClient.GetAttendanceCodeById(request);

            _logger.LogInformation($"Found matching attendance code: {code}");
            return attendanceEvent.CurrentUnixTime > code.Unix 
                ? CodeValidity.OutOfTime 
                : CodeValidity.Success;
        }
        catch(Exception ex)
        {
            _logger.LogError(ex.Message);
            _logger.LogError(ex.StackTrace);
            return CodeValidity.Error;
        }
    }
}

public enum CodeValidity
{
    Success,
    OutOfTime,
    Error
}
