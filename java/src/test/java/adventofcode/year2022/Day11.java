package adventofcode.year2022;

import adventofcode.BaseTest;
import lombok.Data;
import org.junit.Ignore;
import org.junit.Test;

import java.util.*;

public class Day11 extends BaseTest {

	public static final int DAY = 11;

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	Long commonDivider=1l;
	@Data
	class Monkey {
		int id;
		List<Long> currentItems = new ArrayList<>();

		String operation;
		int testNumber;
		int testTrueMonkeyId;
		int testFalseMonkeyId;
		long inspectionCount = 0;


		Monkey(String nName, String sStart, String sOperation, String sTest, String sTestTrue, String sTestFalse){
			this.id = Integer.parseInt(""+nName.charAt(7));
			String[] strS = sStart.substring(18).split(", ");
			for (String s: strS){
				currentItems.add(Long.parseLong(s));
			}
			this.operation = sOperation.substring(17);
			this.testNumber = Integer.parseInt( sTest.substring(21));
			this.testTrueMonkeyId = Integer.parseInt( sTestTrue.substring(29));
			this.testFalseMonkeyId = Integer.parseInt( sTestFalse.substring(30));
		}

		public void processRound(Map<Integer, Monkey> monkeys, boolean gold){
			for (Long current : currentItems){
				Long currWorry = getNewWorry(current);
				if (gold){
					currWorry = currWorry % commonDivider;
				}else{
					currWorry = currWorry / 3;
				}
				if (currWorry%testNumber == 0){
					monkeys.get(testTrueMonkeyId).currentItems.add(currWorry);
				}else{
					monkeys.get(testFalseMonkeyId).currentItems.add(currWorry);
				}
				this.inspectionCount++;
			}
			this.currentItems.clear();
		}


		public Long getNewWorry(Long current) {
			String last = operation.split(" ")[operation.split(" ").length-1];
			Long number = last.equals("old") ? current : Long.parseLong(last);
			return operation.contains("+") ?  current + number : current*number;
		}
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		data.add("");
		Map<Integer, Monkey> monkeys = new HashMap<>();
		for (int i=0; i<data.size(); i=i+7){
			Monkey current = new Monkey(data.get(i), data.get(i+1), data.get(i+2), data.get(i+3), data.get(i+4), data.get(i+5));
			monkeys.put(current.id, current);
			commonDivider = commonDivider*current.getTestNumber();
		}

		for (int j=0; j<10000; j++){
			for (int i=0; i< monkeys.size(); i++){
				monkeys.get(i).processRound(monkeys, true);
			}
		}
		List<Long> inspections = new ArrayList<>();
		for (Map.Entry<Integer, Monkey> entry : monkeys.entrySet()) {
			inspections.add(entry.getValue().inspectionCount);
		}
		Collections.sort(inspections);
		System.out.println(inspections.get(inspections.size()-1) *inspections.get(inspections.size()-2));

	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		data.add("");
		Map<Integer, Monkey> monkeys = new HashMap<>();
		for (int i=0; i<data.size(); i=i+7){
			Monkey current = new Monkey(data.get(i), data.get(i+1), data.get(i+2), data.get(i+3), data.get(i+4), data.get(i+5));
			monkeys.put(current.id, current);
		}

		for (int j=0; j<20; j++){
			for (int i=0; i< monkeys.size(); i++){
				monkeys.get(i).processRound(monkeys, false);
			}
		}
		List<Long> inspections = new ArrayList<>();
		for (Map.Entry<Integer, Monkey> entry : monkeys.entrySet()) {
			inspections.add(entry.getValue().inspectionCount);
		}
		Collections.sort(inspections);
		System.out.println(inspections.get(inspections.size()-1) *inspections.get(inspections.size()-2));
	}

}
