class Solution {

    /**
    * @param String $s
    * @return Integer
    */
    function findTheLongestBalancedSubstring($s) {
        $lastChar = "";
        $ans = 0;
        $m = 0;
        $n = 0;
        for($i = 0; $i < strlen($s); $i++) {
            if ($lastChar == "1" && $s[$i] == "0") {
                $ans =max($ans, $n-abs($m));
                $m = 0;
                $n =0;
            }
            if ($s[$i] == "0") {
                $m++;
            } else {
                $m--;
            }
            $n++;
            $lastChar = $s[$i];
        } 
        $ans =max($ans, $n-abs($m));
        return $ans;
    }
}