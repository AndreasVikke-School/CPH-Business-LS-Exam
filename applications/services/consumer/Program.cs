using Grpc.Net.Client;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;

IHostBuilder host = Host.CreateDefaultBuilder(args)
    .ConfigureServices((hostContext, services) =>
    {
        services.AddOptions<KafkaSettings>().BindConfiguration("Kafka");
        services.AddSingleton<KafkaConsumer>();
        services.AddSingleton<RedisConsumer>();
        services.AddSingleton<PostgresConsumer>();

        RedisSettings redisSettings = hostContext.Configuration.GetSection("Redis").Get<RedisSettings>();
        services.AddSingleton(services => new AttendanceCodeProto.AttendanceCodeProtoClient(GrpcChannel.ForAddress(redisSettings.ServiceUrl)));

        PostgresSettings postgresSettings = hostContext.Configuration.GetSection("Postgres").Get<PostgresSettings>();
        services.AddSingleton(services => new CheckInService.CheckInServiceClient(GrpcChannel.ForAddress(postgresSettings.ServiceUrl)));

        services.AddHostedService<ConsumerService>();
    });

host.Build().Run();

public class ConsumerService : BackgroundService
{
    private readonly KafkaConsumer _kafkaConsumer;
    private readonly RedisConsumer _redisConsumer;
    private readonly PostgresConsumer _postgresConsumer;

    public ConsumerService(KafkaConsumer kafkaConsumer, RedisConsumer redisConsumer, PostgresConsumer postgresConsumer)
    {
        _kafkaConsumer = kafkaConsumer;
        _redisConsumer = redisConsumer;
        _postgresConsumer = postgresConsumer;
    }

    protected override async Task ExecuteAsync(CancellationToken stoppingToken)
        => await Task.Run(async () 
            => {
                while (!stoppingToken.IsCancellationRequested)
                {
                    var attendance = _kafkaConsumer.Poll(stoppingToken);
                    var codeValidity = _redisConsumer.CheckCodeValidity(attendance);
                    await _postgresConsumer.SaveAttendanceEvent(attendance, codeValidity);
                }
            }, stoppingToken);
}