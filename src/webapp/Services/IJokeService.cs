using System.Threading.Tasks;
using ChuckJokes.WebApp.Models;

namespace ChuckJokes.WebApp.Services
{
    public interface IJokeService
    {
        Task<Joke> GetRandomJoke();
    }
}