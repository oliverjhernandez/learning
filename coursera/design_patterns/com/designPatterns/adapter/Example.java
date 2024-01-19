// public interface IWebRequester {
//   public int request(Object);
// }
//
// public class WebAdapter implements IWebRequester {
//   private WebService service;
//
//   public void connect(WebService currentService) {
//     this.service = currentService;
//   }
//
//   public int request(Object request) {
//     Json result = this.toJson(request);
//     Json response = service.request(result);
//     if (response != null) {
//       return 200;
//     }
//     return 500;
//   }
//
//   private Json toJson(Object input) {}
// }
//
// Ungraded Assignment
// public class Main {
//   public static void main(String args[]) {
//     String webHost = "Host: https://google.com";
//     WebService service = new WebService(webHost);
//     WebAdapter adapter = new WebAdapter();
//     adapter.connect(service);
//     WebClient client = new WebClient(adapter);
//     client.doWork();
//   }
// }
