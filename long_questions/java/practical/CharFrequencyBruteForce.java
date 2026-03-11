package long_questions.java.practical;

public class CharFrequencyBruteForce {
    public static void main(String[] args) {
        String str = "programming";
        boolean[] counted=new boolean[str.length()];


        for(int i=0;i<str.length();i++){
            if(!counted[i]){
                char ch=str.charAt(i);
                int count=1;

                for(int j=i+1;i<str.length();j++){
                    if(ch==str.charAt(j)){
                        count++;
                        counted[j]=true;
                    }
                }

                System.out.println(ch + "=" + count);
                counted[i] = true;
            }
        }

    }
}


