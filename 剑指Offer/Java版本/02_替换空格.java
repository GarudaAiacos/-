public class Solution {
    public String replaceSpace(StringBuffer str) {
    	if(str == null) {
            return null;
        }
        
        StringBuffer newStr = new StringBuffer();
        for(int i=0; i<str.length(); i++) {
            if(str.charAt(i) == ' '){
                newStr.append("%");
                newStr.append("2");
                newStr.append("0");
            } else {
                newStr.append(str.charAt(i));
            }
        }
        return newStr.toString();
    }
}