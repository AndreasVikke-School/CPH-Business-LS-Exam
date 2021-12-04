using Geolocation;
using Google.Protobuf.WellKnownTypes;

namespace LSExam.Services;

public class RedisConsumer
{
    private readonly ILogger<RedisConsumer> _logger;
    private readonly AttendanceCodeProto.AttendanceCodeProtoClient _redisClient;

    public RedisConsumer(ILogger<RedisConsumer> logger, AttendanceCodeProto.AttendanceCodeProtoClient redisClient)
    {
        _logger = logger;
        _redisClient = redisClient;
    }

    public CodeValidity CheckCodeValidity(AttendanceEvent attendanceEvent)
    {
        long attandanceCode = attendanceEvent.AttendanceCode;
        Int64Value request = new() { Value = attandanceCode };
        try
        {
            AttendanceCode code = _redisClient.GetAttendanceCodeById(request);

            if (code is { Code: -1, Unix: -1 })
            {
                _logger.LogInformation($"Could not find any attendance code for: {attendanceEvent}");
                return CodeValidity.NotFound;
            }

            _logger.LogInformation($"Found matching attendance code: {code}");

            double distanceInMeters = GeoCalculator.GetDistance(code.Lat, code.Long, attendanceEvent.Latitude, attendanceEvent.Longitude, distanceUnit: DistanceUnit.Meters);
            if (distanceInMeters > 100d)
                return CodeValidity.OutOfRange;

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