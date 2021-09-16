import java.util.Arrays;

public class WildcardMatch {
 
    // Function that matches input str with
    // given wildcard pattern
    static boolean strmatch(String str, String pattern, int n, int m) {
        
        // empty pattern can only match with empty string
        if (m == 0) {
            return (n == 0);
        }
 
        boolean[][] lookup = new boolean[n + 1][m + 1];
        for (int i = 0; i < n + 1; i++) {
            Arrays.fill(lookup[i], false);
        }
 
        lookup[0][0] = true;
        for (int j = 1; j <= m; j++) {
            if (pattern.charAt(j - 1) == '*') { // Only '*' can match with empty string
                lookup[0][j] = lookup[0][j - 1];
            }
        }

        // fill the table in bottom-up fashion
        for (int i = 1; i <= n; i++) {

            for (int j = 1; j <= m; j++) {
                // Two cases if we see a '*'
                // a) We ignore '*'' character and move
                //    to next  character in the pattern,
                //     i.e., '*' indicates an empty
                //     sequence.
                // b) '*' character matches with ith
                //     character in input
                if (pattern.charAt(j - 1) == '*') {
                    lookup[i][j] = lookup[i][j - 1] || lookup[i - 1][j];
                }
                // Current characters are considered as
                // matching in two cases
                // (a) current character of pattern is '?'
                // (b) characters actually match
                else if (pattern.charAt(j - 1) == '?' || str.charAt(i - 1) == pattern.charAt(j - 1)) {
                    lookup[i][j] = lookup[i - 1][j - 1];
                } else {
                    lookup[i][j] = false;
                }  
            }
        }
 
        return lookup[n][m];
    }
 
   
    // Driver code
    public static void main(String args[]) {

        String str = "baaabab";
        String pattern = "*****ba*****ab";
        // String pattern = "aa?ab";
        // String pattern = "b*b";
        // String pattern = "a*a";
        // String pattern = "baaabab";
        // String pattern = "?baaabab";
        // String pattern = "*baaaba*";
 
        if (strmatch(str, pattern, str.length(), pattern.length()))
            System.out.println("Yes");
        else
            System.out.println("No");
    }

}