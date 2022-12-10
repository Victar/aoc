package adventofcode.year2022;

import adventofcode.BaseTest;
import org.apache.commons.lang3.StringUtils;
import org.junit.Ignore;
import org.junit.Test;

import java.util.ArrayList;
import java.util.List;

public class Day10 extends BaseTest {

	public static final int DAY = 10;
	int x = 1;
	int s = 0;
	int cycle = 0;
	int size = 40;
	List<StringBuilder> image = new ArrayList<>();

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runSilverAndGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		for (int i = 0; i < 6; i++) {
			image.add(new StringBuilder(StringUtils.leftPad(StringUtils.EMPTY, 40, '.')));
		}
		for (final String input : data) {
			if (input.startsWith("noop")) {
				doCycle();
			}
			if (input.startsWith("addx")) {
				doCycle();
				doCycle();
				x = x + Integer.parseInt(input.split(" ")[1]);
			}
		}
		System.out.println(s);
		for (StringBuilder s : image) {
			System.out.println(s.toString());
		}
	}

	public void doCycle() {
		//Gold
		int raw = cycle / size;
		int pos = cycle % size;
		if (Math.abs(x - pos) < 2) {
			image.get(raw).setCharAt(pos, '#');
		}
		//Silver
		cycle++;
		if ((cycle + 20) % size == 0) {
			s = s + (cycle) * x;
		}
	}

}
