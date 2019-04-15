import java.io.*;
import java.net.*;

public class YourServer
{    
    public static void main(String[] args ) 
    {
        try {    
            ServerSocket socket = new ServerSocket(8765);

            Socket insocket = socket.accept( );

            BufferedReader in = new BufferedReader (new 
                InputStreamReader(insocket.getInputStream()));
            PrintWriter out = new PrintWriter (insocket.getOutputStream(), 
                true);

            String instring = in.readLine();
            out.println("The server got this: " + instring);
            insocket.close();
        }
        catch (Exception e) {} 
     } 
}
