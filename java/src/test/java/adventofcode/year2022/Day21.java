package adventofcode.year2022;

import adventofcode.BaseTest;
import lombok.Data;
import org.junit.Ignore;
import org.junit.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

public class Day21 extends BaseTest {

	public static final int DAY = 21;

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		Map<String,Monkey> monkeys= new HashMap<>();
		Monkey root = null;
		for (final String input : data) {
			Monkey current = new Monkey(input);
			monkeys.put(current.name, current);
		}
		System.out.println(monkeys.get("root").getNumber(monkeys, new HashMap<>()));
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		Map<String,Monkey> monkeys= new HashMap<>();
		Monkey root = null;
		Monkey rootLeft = null;
		Monkey rootRight = null;
		String meName = "humn";
		for (final String input : data) {
			Monkey current = new Monkey(input);
			monkeys.put(current.name, current);
		}
		HashMap<String, Long> DP = new HashMap<>();
		root = monkeys.get("root");
		rootLeft = monkeys.get(root.monkeyLeft);
		rootRight = monkeys.get(root.monkeyRight);
		long i  = 3352886100000L;
//		long i  = 52716091087786
//		52716091090000 leftNumber -730962523154826  52716091087786 summ -783678614242612
// 		62716091090000 leftNumber -889720167626738  52716091087786 summ -942436258714524
//		62716091110000 leftNumber -889720167944250  52716091087786 summ -942436259032036
//		42716091090000 leftNumber -572204878682870  52716091087786 summ -624920969770656
//		2990000005787 leftNumber 58477185777786  52716091087786    diff    5761094690000
//		3000000011175 leftNumber 58318428047786  52716091087786    diff    5602336960000
//		3200000008691 leftNumber 55143275197786  52716091087786    diff    2427184110000
//		3500000002761 leftNumber 50380545957786  52716091087786    diff   -2335545130000
//		3354999906694 leftNumber 52682533327786  52716091087786 diff -33557760000
//		3354999909211 leftNumber 52682533287786  52716091087786 diff -33557800000
//		3353999902563 leftNumber 52698409157786  52716091087786 diff -17681930000
//		3352999901583 leftNumber 52714284937786  52716091087786 diff  -1806150000
//		3356000009083 leftNumber 52666655937786  52716091087786 diff -49435150000
//		3352500000615 leftNumber 52722221247786  52716091087786 diff   6130160000
		while(true){
			i++;
			DP.clear();
			DP.put(meName, i);
			long leftNumber = rootLeft.getNumber(monkeys, DP);
			long rightNumber = rootRight.getNumber(monkeys, DP);
			if ((leftNumber - rightNumber)%10000==0){
				System.out.println(i + " leftNumber " + leftNumber + "  " + rightNumber + " diff " + (leftNumber - rightNumber));
			}

			if (leftNumber == rightNumber){
				System.out.println(i);
				break;
			}
		}
//		System.out.println(monkeys.get("root").getNumber(monkeys, new HashMap<>()));
	}

	@Data
	class Monkey{
		String name;
		String operation;
		boolean isSimple;
		long simpleNumber;
		String monkeyLeft;
		String monkeyRight;
		String monkeyOperation;


		public Monkey(String input){
			String[] arr = input.split(": ");
			this.name = arr[0];
			this.operation = arr[1];
			try{
				this.simpleNumber = Integer.parseInt(arr[1]);
				this.isSimple = true;
			}catch(Exception ex){}
			if (!isSimple){
				String[] ops = this.operation.split(" ");
				monkeyLeft = ops[0];
				monkeyOperation = ops[1];
				monkeyRight = ops[2];
			}
		}

		public long getNumber(Map<String, Monkey> monkeys, Map<String, Long> DP){
			if (DP.containsKey(this.name)){
				return DP.get(this.name);
			}
			if (isSimple){
				DP.put(this.name, this.simpleNumber);
				return simpleNumber;
			}
			long leftNumber = monkeys.get(monkeyLeft).getNumber(monkeys, DP);
			long rightNumber = monkeys.get(monkeyRight).getNumber(monkeys, DP);
			long result = 0;
			if ("+".equalsIgnoreCase(monkeyOperation)){
				return leftNumber + rightNumber;
			}
			if ("-".equalsIgnoreCase(monkeyOperation)){
				return leftNumber - rightNumber;
			}
			if ("*".equalsIgnoreCase(monkeyOperation)){
				return leftNumber * rightNumber;
			}
			if ("/".equalsIgnoreCase(monkeyOperation)){
				return leftNumber / rightNumber;
			}

			DP.put(this.name, result);
			return 0;
		}
	}

}
