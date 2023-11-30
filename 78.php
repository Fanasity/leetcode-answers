class Solution {
    /**
    * @param Integer[] $nums
    * @return Integer[][]
    */
    function subsets($nums) {
        $len = count($nums);
        $answer = [[], $nums];
        for($i = 1; $i < $len; $i++ ){
            // 取出包含指定个数的子集
            $answer = array_merge($answer, $this->getSets([], $nums, $i));
        }
        return $answer;
    }

    /**
    * @param Integer[] $com 已取出的元素数组
    * @param Integer[] $array 
    * @param Integer $count 还需要取出几个元素
    * @return Integer[][] 
    */
    function getSets($com, $array, $count){
        $sets = [];
        foreach($array as $k => $v){
            $set = $com;
            $set[] = $v;
            // 是否是取最后一个元素
            if($count != 1){
                $newArray = $array;
                unset($newArray[$k]);
                // 递归
                $sets = array_merge($sets, $this->getSets($set, $newArray, $count - 1));
            }else{
                $sets[] = $set;  
            }
            //已使用的元素释放掉
            unset($array[$k]);
        }
        return $sets;
    }
}