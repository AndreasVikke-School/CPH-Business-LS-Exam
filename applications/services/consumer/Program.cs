using LSExam.Configs;
using LSExam.Services;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;

IHostBuilder host = Host.CreateDefaultBuilder(args)
    .ConfigureServices((hostContext, services) =>
    {
        services.AddOptions<KafkaSettings>().BindConfiguration("Kafka");
        services.AddSingleton<KafkaConsumer>();
        services.AddHostedService<ConsumerService>();
    });

host.Build().Run();

public class ConsumerService : BackgroundService
{
    private readonly KafkaConsumer _kafkaConsumer;

    public ConsumerService(KafkaConsumer kafkaConsumer)
    {
        _kafkaConsumer = kafkaConsumer;
    }

    protected override Task ExecuteAsync(CancellationToken stoppingToken)
        => Task.Run(() 
            => {
                while (!stoppingToken.IsCancellationRequested)
                {
                    var attendance = _kafkaConsumer.Poll(stoppingToken);
                    Console.WriteLine(attendance);
                }
            }, stoppingToken);
}