class Solution {

    /**
    * @param Integer[] $nums
    * @param Integer $target
    * @return Integer[]
    */
    function twoSum($nums, $target) {
        foreach($nums as $k => $v){
            $res = array_search(($target-$v), $nums);
            if($res !== false && $res != $k){
                return [$k, $res];
            }
        }
        return [];
    }
}